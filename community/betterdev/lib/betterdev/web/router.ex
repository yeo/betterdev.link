defmodule Betterdev.Web.Router do
  use Betterdev.Web, :router

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_flash
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  pipeline :api do
    plug :accepts, ["json"]
    plug Joken.Plug,
      verify: &Betterdev.JWTHelpers.verify/0,
      on_error: &Betterdev.JWTHelpers.error/2
  end

  scope "/", Betterdev.Web do
    pipe_through :browser # Use the default browser stack

    get "/", PageController, :index
  end

  # Other scopes may use custom stacks.
  scope "/api", Betterdev.Web do
    pipe_through :api

    get "/status", StatusController, :index
    resources "/links", LinkController, except: [:new, :edit]
  end
end
