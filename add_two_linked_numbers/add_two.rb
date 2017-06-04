# Definition for singly-linked list.
class ListNode
  attr_accessor :val, :next
  def initialize(val)
    @val = val
    @next = nil
  end
end

# @param {ListNode} l1
# @param {ListNode} l2
# @return {ListNode}
def add_two_numbers(l1, l2)
  num1 = get_vals(l1).reverse
  num2 = get_vals(l2).reverse
  res = num1.join.to_i + num2.join.to_i
  nodes = []
  res.to_s.split('').reverse_each { |s| nodes << ListNode.new(s.to_i) }
  nodes.each_with_index {|e, i| e.next = nodes[i+1]}
  nodes[0]
end

def get_vals(node)
  nums = []
  until node.nil?
    nums << node.val
    node = node.next
  end
  nums
end

n1 = ListNode.new(1)
n2 = ListNode.new(8)
n1.next = n2

l1 = ListNode.new(0)

add_two_numbers(n1, l1)
