local mymod =require("mymod")

function init()
    global_id = 1
    global_name = "test"
end

function newReader()
    r = reader.new(global_id,global_name,0)
end

-- 连续执行三次
function read(book)
    r:read(book)
    mymod.eat("面包")
    mymod.drink("雪碧")
end

function finish()
    mymod.record(r)
end