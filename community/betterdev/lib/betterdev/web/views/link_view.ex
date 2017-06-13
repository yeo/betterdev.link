defmodule Betterdev.Web.LinkView do
  use Betterdev.Web, :view
  alias Betterdev.Web.LinkView

  def render("index.json", %{links: links}) do
    %{data: render_many(links, LinkView, "link.json")}
  end

  def render("show.json", %{link: link}) do
    %{data: render_one(link, LinkView, "link.json")}
  end

  def render("link.json", %{link: link}) do
    %{id: link.id,
      title: link.title,
      uri: link.uri}
  end
end
