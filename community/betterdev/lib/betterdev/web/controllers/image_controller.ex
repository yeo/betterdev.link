defmodule Betterdev.Web.ImageController do
  use Betterdev.Web, :controller

  def index(conn, params}) do
    url = params["url"]
    case url do
      nil ->
        html conn, ""

      url ->
        HTTPoison.start
        hackney = [follow_redirect: true]
        response = HTTPoison.get!(url, [], [ hackney: hackney ])

        html conn, response.body
    end
  end
end
