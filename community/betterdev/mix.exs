defmodule Betterdev.Mixfile do
  use Mix.Project

  def project do
    [app: :betterdev,
     version: "0.0.1",
     elixir: "~> 1.4",
     elixirc_paths: elixirc_paths(Mix.env),
     compilers: [:phoenix, :gettext] ++ Mix.compilers,
     start_permanent: Mix.env == :prod,
     aliases: aliases(),
     deps: deps()]
  end

  # Configuration for the OTP application.
  #
  # Type `mix help compile.app` for more information.
  def application do
    [mod: {Betterdev.Application, []},
      #extra_applications: [:logger, :runtime_tools, :exbot, :algolia, :exq, :exq_ui]]
     extra_applications: [:new_relic, :logger, :runtime_tools, :exbot, :tirexs]]
  end

  # Specifies which paths to compile per environment.
  defp elixirc_paths(:test), do: ["lib", "test/support"]
  defp elixirc_paths(_),     do: ["lib"]

  # Specifies your project dependencies.
  #
  # Type `mix help deps` for examples and options.
  defp deps do
    [{:phoenix, "~> 1.3.0-rc"},
     {:phoenix_pubsub, "~> 1.0"},
     {:phoenix_ecto, "~> 3.2"},
     {:ecto, "~> 2.0"},
     {:mariaex, ">= 0.0.0"},
     {:phoenix_html, "~> 2.6"},
     {:phoenix_live_reload, "~> 1.0", only: :dev},
     {:gettext, "~> 0.11"},
     {:joken, "~> 1.1"},
     {:exbot, ">= 0.0.1"},
     {:scrape, "~> 2.0.0"},
     {:kerosene, "~> 0.7.0"},
     #{:exq, "~> 0.9.0"},
     #{:exq_ui, "~> 0.9.0"},
     {:tirexs, "~> 0.8"},
     {:new_relic, git: "https://github.com/runtastic/newrelic-elixir.git"},
     {:timex, "~> 3.0"},
     {:slack, "~> 0.11.0"},
     {:cowboy, "~> 1.0"}]
  end

  # Aliases are shortcuts or tasks specific to the current project.
  # For example, to create, migrate and run the seeds file at once:
  #
  #     $ mix ecto.setup
  #
  # See the documentation for `Mix` for more info on aliases.
  defp aliases do
    ["ecto.setup": ["ecto.create", "ecto.migrate", "run priv/repo/seeds.exs"],
     "ecto.reset": ["ecto.drop", "ecto.setup"],
     "test": ["ecto.create --quiet", "ecto.migrate", "test"]]
  end
end
