box.cfg {
    listen = 3301
}

local users = box.schema.create_space('users', { format = {
    { name = 'id', type = 'unsigned' },
    { name = 'first_name', type = 'string' },
    { name = 'last_name', type = 'string' },
    { name = 'age', type = 'unsigned' },
    { name = 'sex', type = 'string' },
    { name = 'interests', type = 'string' },
    { name = 'password', type = 'string' },
    { name = 'created_at', type = 'unsigned' }
}, if_not_exists = true })

users:create_index('pk', { parts = { 'id' }, if_not_exists = true })
-- users:create_index('first_name_last_name', {
--     parts = {
--         {'first_name'},
--         {'last_name'}
--     },
--         unique = false,
--         if_not_exists = true
-- })
box.space.users:create_index('last_name', {
    parts = {'last_name'},
    unique = false,
    if_not_exists = true
})
box.space.users:create_index('first_name', {
    parts = {'first_name'},
    unique = false,
    if_not_exists = true
})

-- users:create_index('last_name', {parts = {'last_name'}, if_not_exists = true })

box.schema.sequence.create('users_id', { if_not_exists = true })
-- box.schema.func.drop('create_user', {if_not_exists = true})
box.schema.func.create('create_user', {if_not_exists = true})
box.schema.func.create('get_user', {if_not_exists = true})
box.schema.func.create('find_user', {if_not_exists = true})


box.func.find_user = function(name)
    local result = {}
    local search = name .. '%'
    for _, user in box.space.users.first_name:pairs(search) do
        table.insert(result, user)
    end
    for _, user in box.space.users.first_name:pairs(search) do
        table.insert(result, user)
    end
    return result
end

box.func.create_user = function(first_name, last_name, age, sex, interests, password)
    local id = box.sequence.users_id:next()
    local created_at = os.time()
    local user = {id, first_name, last_name, age, sex, interests, password, created_at}
--     box.space.users:insert(user)
--     return user
    local last_id = id - 1
        -- Loop until we find the next available id
        while box.space.users:get(id) ~= nil do
            last_id = id
            id = box.sequence.users_id:next()
            user[1] = id
        end
        -- Update sequence if necessary
        if last_id ~= id - 1 then
            box.sequence.users_id:set(id)
        end
        box.space.users:insert(user)
        return user
end

box.func.get_user = function(user_id)
    local result = box.space.users:select{user_id}
        if #result == 0 then
            return nil
        else
            return result[1]
        end
end


--
-- function multi_insert_users(data)
--     local tuples = {}
--     for i, row in ipairs(data) do
--         local tuple = {
--             row.id,
--             row.first_name,
--             row.last_name,
--             row.age,
--             row.sex,
--             row.interests,
--             row.password,
--             os.time()
--         }
--         table.insert(tuples, tuple)
--     end
--     box.space.users:insert(tuples)
-- end

