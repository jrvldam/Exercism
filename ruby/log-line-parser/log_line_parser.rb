class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    return @line.gsub(/\[(INFO|WARNING|ERROR)\]:/, '').strip
  end

  def log_level
    return @line.slice(1, @line.index(':') - 2).downcase
  end

  def reformat
    return "#{message} (#{log_level})"
  end
end
