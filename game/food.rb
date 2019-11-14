class Food
  attr_reader :board
  attr_reader :position

  def initialize(board, x, y)
    @board = board
    @position = board[x, y]

    position.food = self
  end

  def inspect
    "<Food x=#{position.x} y=#{position.y}>"
  end
end
