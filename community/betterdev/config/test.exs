use Mix.Config

# We don't run a server during test. If one is required,
# you can enable the server option below.
config :betterdev, Betterdev.Web.Endpoint,
  http: [port: 4001],
  server: false

# Print only warnings and errors during test
config :logger, level: :warn

# Configure your database
config :betterdev, Betterdev.Repo,
  adapter: Ecto.Adapters.MySQL,
  username: "root",
  password: "",
  database: "betterdev_test",
  hostname: "localhost",
  pool: Ecto.Adapters.SQL.Sandbox
