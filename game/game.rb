require_relative 'board'

class Game
  attr_reader :data

  MOVES = %w[up down left right].map(&:freeze).freeze

  def initialize(data)
    @data = data

    p turn
    p board.food
    p board.snakes
    p board.snakes.map(&:segments)
  end

  %w[game turn board you].each do |key|
    define_method("#{key}_data") { data[key] }
  end

  alias turn turn_data

  def start
    # Add new game to db
  end

  def move
    { move: decide_move }
  end

  def end
    # End game in db
  end

  def id
    @id ||= game_data['id']
  end

  def you
    @you ||= board.snakes.find { |snake| snake.id == you_data['id'] }
  end

  def board
    @board ||= Board.new(self, board_data)
  end

  def decide_move
    moves = livable_moves
    p moves

    unless board.food.empty?
      food_distance = you.head.position.distance(nearest_food.position)
      food_moves = moves_for(moves) { |x, y| board[x, y].distance(nearest_food.position) < food_distance }
      return random_move(moves) if food_moves.empty?
      moves = food_moves
    end

    p moves
    random_move(moves)
  end

  def livable_moves
    moves = valid_moves
    p moves

    moves_for(moves) { |x, y| board[x, y].segments.all? { |segment| segment.head? && segment.snake.killable? } }
  end

  def valid_moves
    moves_for { |x, y| board.valid_xy(x, y) }
  end

  def moves_for(moves = MOVES, &block)
    moves.select { |move| block.call(*you.head.position.send(move)) }
  end

  def random_move(moves = MOVES)
    moves.sample
  end

  def nearest_food
    @nearest_food ||= board.food.min_by { |f| f.position.distance(you.head.position) }
  end
end
