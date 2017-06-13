defmodule Betterdev.CommunityTest do
  use Betterdev.DataCase

  alias Betterdev.Community

  describe "links" do
    alias Betterdev.Community.Link

    @valid_attrs %{title: "some title", uri: "some uri"}
    @update_attrs %{title: "some updated title", uri: "some updated uri"}
    @invalid_attrs %{title: nil, uri: nil}

    def link_fixture(attrs \\ %{}) do
      {:ok, link} =
        attrs
        |> Enum.into(@valid_attrs)
        |> Community.create_link()

      link
    end

    test "list_links/0 returns all links" do
      link = link_fixture()
      assert Community.list_links() == [link]
    end

    test "get_link!/1 returns the link with given id" do
      link = link_fixture()
      assert Community.get_link!(link.id) == link
    end

    test "create_link/1 with valid data creates a link" do
      assert {:ok, %Link{} = link} = Community.create_link(@valid_attrs)
      assert link.title == "some title"
      assert link.uri == "some uri"
    end

    test "create_link/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Community.create_link(@invalid_attrs)
    end

    test "update_link/2 with valid data updates the link" do
      link = link_fixture()
      assert {:ok, link} = Community.update_link(link, @update_attrs)
      assert %Link{} = link
      assert link.title == "some updated title"
      assert link.uri == "some updated uri"
    end

    test "update_link/2 with invalid data returns error changeset" do
      link = link_fixture()
      assert {:error, %Ecto.Changeset{}} = Community.update_link(link, @invalid_attrs)
      assert link == Community.get_link!(link.id)
    end

    test "delete_link/1 deletes the link" do
      link = link_fixture()
      assert {:ok, %Link{}} = Community.delete_link(link)
      assert_raise Ecto.NoResultsError, fn -> Community.get_link!(link.id) end
    end

    test "change_link/1 returns a link changeset" do
      link = link_fixture()
      assert %Ecto.Changeset{} = Community.change_link(link)
    end
  end
end
