defmodule Betterdev.StatusView do
  use Betterdev.Web, :view

  def render("status.json", %{status: status}) do
    status
  end
end
