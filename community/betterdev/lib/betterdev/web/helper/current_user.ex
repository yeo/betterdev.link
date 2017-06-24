defmodule Betterdev.Helper.CurrentUser do
  #require Plug.Conn
  use Betterdev.Web, :router

  @doc """
  Detect user from token
  """
  def detect(conn, opts \\ []) do
    claims = Map.get(conn.assigns, :joken_claims)
    user = Accounts.user_from_token(conn.assigns.joken_claims)
    IO.inspect user
    conn |> assign(:current_user, user)
  end
end
