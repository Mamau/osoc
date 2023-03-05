-- Подключаем модуль box
box.cfg {
    listen = 3301,
    log_level = 5,
    wal_mode = 'none',
}

-- Определяем схему
box.once('schema', function()
    box.schema.create_space('users', {
        if_not_exists = true,
    })
    box.space.users:format({
        {name = 'id', type = 'unsigned'},
        {name = 'first_name', type = 'string'},
        {name = 'last_name', type = 'string'},
        {name = 'age', type = 'unsigned'},
        {name = 'sex', type = 'string'},
        {name = 'interests', type = 'string'},
        {name = 'password', type = 'string'},
        {name = 'created_at', type = 'unsigned'},
    })
    box.space.users:create_index('primary', {
        type = 'hash',
        parts = {'id'},
        if_not_exists = true,
    })
end)

-- Определяем процедуру getUsers
function getUsers()
    local users = {}
    for _, tuple in box.space.users:pairs() do
        table.insert(users, {
            id = tuple[1],
            first_name = tuple[2],
            last_name = tuple[3],
            age = tuple[4],
            sex = tuple[5],
            interests = tuple[6],
            password = tuple[7],
            created_at = tuple[8],
        })
    end
    return users
end

-- Сохраняем процедуру getUsers
box.schema.func.create('getUsers', {
    if_not_exists = true,
    body = function()
        return getUsers()
    end
})