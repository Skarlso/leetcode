# @param {String} s
# @return {Integer}
def length_of_longest_substring(s)
  get_repeating_subs(s)
end


def get_repeating_subs(s)
  maxi = 0
  i = 0
  while i < s.length
    j = i
    while j < s.length + 1
      sub = s[i..j]
      uniq, index = unique_with_index(sub)
      if uniq
        maxi = sub.length if maxi < sub.length
      else
        i += index
        break
      end
      j += 1
    end
    i += 1
  end
  maxi
end

# Returns if the string is unique.
# If not, it also returns the index of the character which was the first duplicate.
def unique_with_index(s)
  # puts "s: #{s}"
  return [true, nil] if s.chars.to_a.uniq.length == s.length
  s.chars.to_a.each_with_index do |e, i|
    # puts "e: #{e}; i: #{i}"
    return [false, i] if s.count(e) > 1
  end
end

p length_of_longest_substring('iuonzyzlclfudgrr')
