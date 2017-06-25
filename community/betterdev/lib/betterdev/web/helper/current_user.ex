defmodule Betterdev.Helper.CurrentUser do
  import Plug.Conn
  alias Betterdev.Accounts

  def init(opts \\ %{}) do
  end

  def call(conn, _params) do
    detect(conn)
  end

  @doc """
  Detect user from token
  """
  def detect(conn, opts \\ []) do
    claims = Map.get(conn.assigns, :joken_claims)
    case claims do
      nil -> conn
      _ ->
        user = Accounts.user_from_token(conn.assigns.joken_claims)
        conn |> assign(:current_user, user)
    end
  end
end
