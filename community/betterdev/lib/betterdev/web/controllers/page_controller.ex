defmodule Betterdev.Web.PageController do
  use Betterdev.Web, :controller

  alias Betterdev.Community

  def index(conn, params) do
    {postings, kerosene} = Community.list_links(params)
    assigns = [
      postings: postings,
      kerosene: kerosene
    ]
  render conn, "index.html", assigns
  end
end
