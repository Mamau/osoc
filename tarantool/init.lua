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

box.schema.func.create('create_user', {if_not_exists = true})
box.schema.user.grant('guest', 'execute', 'function', 'create_user')

box.func.create_user = function(first_name, last_name, age, sex, interests, password)
    local id = box.sequence.users_id:next()
    local created_at = os.time()
    local user = {id, first_name, last_name, age, sex, interests, password, created_at}
    box.space.users:insert(user)
    return user
end

--box.schema.func.create('create_user', {if_not_exists = true})
--
--box.schema.func.create_user:replace(function(first_name, last_name, age, sex, interests, password)
--    local id = box.sequence.users_id:next()
--    local created_at = os.time()
--    local user = {id, first_name, last_name, age, sex, interests, password, created_at}
--    box.space.users:insert(user)
--    return user
--end)
