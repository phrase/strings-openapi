require "json"

run lambda { |env|
  path = env["PATH_INFO"]

  case path
  when "/ping"
    [200, { "Content-Type" => "application/json" }, [JSON.dump({ message: "pong" })]]
  when "/status/500"
    [500, { "Content-Type" => "application/json" }, [JSON.dump({ error: "boom" })]]
  else
    [404, { "Content-Type" => "application/json" }, [JSON.dump({ error: "not found" })]]
  end
}
