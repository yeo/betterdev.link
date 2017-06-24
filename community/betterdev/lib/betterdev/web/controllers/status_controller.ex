defmodule Betterdev.Web.StatusController do
  use Betterdev.Web, :controller

  def index(conn, _params) do
    render(conn, "status.json", status: conn.assigns.current_user.email)
  end
end
