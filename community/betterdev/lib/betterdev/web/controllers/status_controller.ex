defmodule Betterdev.Web.StatusController do
  use Betterdev.Web, :controller

  def index(conn, _params) do
    status = %{
      success: true
    }
    IO.inspect conn.assigns.joken_claims
    render(conn, "status.json", status: status)
  end
end
