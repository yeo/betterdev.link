defmodule Betterdev.Web.PageController do
  use Betterdev.Web, :controller

  def index(conn, _params) do
    render conn, "index.html"
  end
end
