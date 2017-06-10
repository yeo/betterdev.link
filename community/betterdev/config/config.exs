# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.
use Mix.Config

# General application configuration
config :betterdev,
  ecto_repos: [Betterdev.Repo]

# Configures the endpoint
config :betterdev, Betterdev.Web.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "Bk0U+ypmzC6sTKW/FXYdWr1FFR6sCTIB0H3hH3fXj29o+S1/ZaK8ifBmCBfz39p6",
  render_errors: [view: Betterdev.Web.ErrorView, accepts: ~w(html json)],
  pubsub: [name: Betterdev.PubSub,
           adapter: Phoenix.PubSub.PG2]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env}.exs"
