defmodule MixexampleTest do
  use ExUnit.Case
  doctest Mixexample

  test "greets the world" do
    assert Mixexample.hello() == :world
  end
end
