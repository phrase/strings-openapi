require "open3"
require "spec_helper"

RSpec.describe "General behavior" do
  it "info prints the version" do
    r = run_cli("info")
    expect(r[:exit_code]).to eq(0)
    expect(r[:stdout]).to include("Phrase Strings client version")
    expect(r[:stderr]).to include("You're running a development version")
  end

  it "handles server 500s cleanly" do
    # Example: pass a flag so your CLI requests /status/500, if it supports it.
    # Otherwise adapt your rack app to respond 500 for the request your CLI already makes.
    r = run_cli("call", "/status/500")
    expect(r[:exit_code]).not_to eq(0)
    expect(r[:stderr]).to match(/500|error/i)
  end
end
