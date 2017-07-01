defmodule Betterdev.Web.Router do
  use Betterdev.Web, :router

  @skip_token_verification %{joken_skip: true}

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
    plug Betterdev.Helper.CurrentUser
  end

  scope "/", Betterdev.Web do
    pipe_through :browser # Use the default browser stack

    get "/", PageController, :index
    get "/img", ImageController, :index
  end

  # Other scopes may use custom stacks.
  scope "/api", Betterdev.Web do
    pipe_through :api

    get "/status", StatusController, :index
    get "/me", MeController, :index

    get  "/links", LinkController, :index, private: @skip_token_verification
    post "/links", LinkController, :create
    get  "/links/:id", LinkController, :show
    #resources "/links", LinkController, except: [:new, :edit]
    resources "/collections", CollectionController, except: [:new, :edit]
  end


	#pipeline :exq do
  #  plug :accepts, ["html"]
  #  plug :fetch_session
  #  plug :fetch_flash
  #  plug :put_secure_browser_headers
  #  plug ExqUi.RouterPlug, namespace: "exq"
  #end

  #scope "/exq", ExqUi do
  #  pipe_through :exq
  #  forward "/", RouterPlug.Router, :index
  #end

end
