defmodule Betterdev.Web.SlackController do
  use Betterdev.Web, :controller

  def index(conn, params) do
    code = params["code"]
    response = Slack.Web.Oauth.access(client_id, client_secret, params["code"], %{redirect_uri: "http://127.0.0.1:4000/bot/slack"})
    IO.inspect response
    html conn, response
  end

  defp client_id, do: Application.get_env(:slack, :client_id)
  defp client_secret, do: Application.get_env(:slack, :client_secret)
end
