defmodule Betterdev.Web.MeController do
  use Betterdev.Web, :controller
  alias Betterdev.Accounts

  def index(conn, _params) do
    user = conn.assigns.current_user
    render(conn, "status.json", status: %{"email" => user.email, "sub" => user.jwt_sub})
  end
end
