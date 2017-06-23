defmodule Betterdev.Web.MeController do
  use Betterdev.Web, :controller

  def index(conn, _params) do
    status = %{
      success: true
    }
    IO.inspect conn
    render(conn, "status.json", status: status)
  end
end
