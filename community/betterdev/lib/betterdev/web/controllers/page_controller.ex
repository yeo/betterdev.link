defmodule Betterdev.Web.PageController do
  use Betterdev.Web, :controller

  alias Betterdev.Community

  def index(conn, _params) do
    assigns =
    [
      postings: Community.list_links
    ]
    render conn, "index.html", assigns
  end
end
