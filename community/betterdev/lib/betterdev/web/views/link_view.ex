defmodule Betterdev.Web.LinkView do
  use Betterdev.Web, :view
  alias Betterdev.Web.LinkView
  import Kerosene.JSON

  def render("index.json", %{links: links, conn: conn, kerosene: kerosene}) do
    %{data: render_many(links, LinkView, "link.json"),  pagination: paginate(conn, kerosene)}
  end

  def render("show.json", %{link: link}) do
    %{data: render_one(link, LinkView, "link.json")}
  end

  def render("link.json", %{link: link}) do
    IO.inspect link.tags

    %{id: link.id,
      title: link.title,
      picture: link.picture || "https://betterdev.link/images/icon-64.png",
      description: link.description,
      tags: render_many(link.tags || [], LinkView, "tag.json"),
      uri: link.uri}
  end

  def render("tag.json", %{link: tag}) do
    %{tag: tag.title, id: tag.id}
  end
end
