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

  describe "collections" do
    alias Betterdev.Community.Collection

    @valid_attrs %{name: "some name", user_id: 42}
    @update_attrs %{name: "some updated name", user_id: 43}
    @invalid_attrs %{name: nil, user_id: nil}

    def collection_fixture(attrs \\ %{}) do
      {:ok, collection} =
        attrs
        |> Enum.into(@valid_attrs)
        |> Community.create_collection()

      collection
    end

    test "list_collections/0 returns all collections" do
      collection = collection_fixture()
      assert Community.list_collections() == [collection]
    end

    test "get_collection!/1 returns the collection with given id" do
      collection = collection_fixture()
      assert Community.get_collection!(collection.id) == collection
    end

    test "create_collection/1 with valid data creates a collection" do
      assert {:ok, %Collection{} = collection} = Community.create_collection(@valid_attrs)
      assert collection.name == "some name"
      assert collection.user_id == 42
    end

    test "create_collection/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Community.create_collection(@invalid_attrs)
    end

    test "update_collection/2 with valid data updates the collection" do
      collection = collection_fixture()
      assert {:ok, collection} = Community.update_collection(collection, @update_attrs)
      assert %Collection{} = collection
      assert collection.name == "some updated name"
      assert collection.user_id == 43
    end

    test "update_collection/2 with invalid data returns error changeset" do
      collection = collection_fixture()
      assert {:error, %Ecto.Changeset{}} = Community.update_collection(collection, @invalid_attrs)
      assert collection == Community.get_collection!(collection.id)
    end

    test "delete_collection/1 deletes the collection" do
      collection = collection_fixture()
      assert {:ok, %Collection{}} = Community.delete_collection(collection)
      assert_raise Ecto.NoResultsError, fn -> Community.get_collection!(collection.id) end
    end

    test "change_collection/1 returns a collection changeset" do
      collection = collection_fixture()
      assert %Ecto.Changeset{} = Community.change_collection(collection)
    end
  end
end
