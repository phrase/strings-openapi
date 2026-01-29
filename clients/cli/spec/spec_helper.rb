require "rspec"
require "open3"
require "socket"
require "timeout"

def free_port
  s = TCPServer.new("127.0.0.1", 0)
  port = s.addr[1]
  s.close
  port
end

RSpec.configure do |config|
  config.before(:suite) do
    @mock_port = free_port
    @mock_url  = "http://127.0.0.1:#{@mock_port}"
    puts "Starting mock server at #{@mock_url}"

    ru_path = File.expand_path("support/mock_server.ru", __dir__)

    # Start rackup in background
    @rack_stdin, @rack_stdout, @rack_stderr, @rack_wait_thr =
      Open3.popen3(
        "bundle", "exec", "rackup",
        "--host", "127.0.0.1",
        "--port", @mock_port.to_s,
        ru_path
      )

    # Wait until the server is accepting connections
    Timeout.timeout(5) do
      loop do
        begin
          TCPSocket.new("127.0.0.1", @mock_port).close
          break
        rescue Errno::ECONNREFUSED, Errno::EHOSTUNREACH
          sleep 0.05
        end
      end
    end

    # Expose to CLI; adapt env var name to whatever your binary uses.
    ENV["BASE_URL"] = @mock_url
  end

  config.after(:suite) do
    ENV.delete("BASE_URL")

    if @rack_wait_thr&.alive?
      Process.kill("TERM", @rack_wait_thr.pid)
      begin
        Timeout.timeout(2) { @rack_wait_thr.join }
      rescue Timeout::Error
        Process.kill("KILL", @rack_wait_thr.pid)
      end
    end

    [@rack_stdin, @rack_stdout, @rack_stderr].compact.each do |io|
      io.close rescue nil
    end
  end
end

# Point this to your actual binary (relative to repo root, or absolute path)
CLI_PATH = File.expand_path("../phrase-cli", __dir__)
# e.g. if it’s built somewhere: File.expand_path("../../target/release/my_cli", __dir__)

def run_cli(*args, env: {})
  full_env = {
    # Whatever your CLI reads to find the server:
    # BASE_URL is set in spec_helper, but you can override per test via env:
    "BASE_URL" => ENV.fetch("BASE_URL")
  }.merge(env)

  stdout, stderr, status = Open3.capture3(full_env, CLI_PATH, *args.map(&:to_s))
  { stdout: stdout, stderr: stderr, status: status, exit_code: status.exitstatus }
end
