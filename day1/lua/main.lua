local input = require("input")

local solution = {}

function solution:compareLists()
	self.result = 0
	for i = 1, #self.list1 do
		if self.list1[i] > self.list2[i] then
			self.result = self.result + (self.list1[i] - self.list2[i])
		else
			self.result = self.result + (self.list2[i] - self.list1[i])
		end
	end

	print(self.result)
end

function solution:splitLists()
	self.list1 = {}
	self.list2 = {}

	-- Split the input into two lists
	for line in input:gmatch("[^\n]+") do
		local num1, num2 = line:match("(%d+)%s+(%d+)")
		table.insert(self.list1, tonumber(num1))
		table.insert(self.list2, tonumber(num2))
	end

	-- Oder both lists
	table.sort(self.list1)
	table.sort(self.list2)

	-- I will consider that both lists have the same size
	self:compareLists()
end

function solution:main()
	self:splitLists()
end

solution:main()
