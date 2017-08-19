var shortid = require('shortid')
var categories = {}

function handleCreate(call, callback) {
    var id = shortid.generate()
    var newCat
    try {
        newCat = JSON.parse(call.request.data)
        newCat.id = id
        categories[id] = newCat
        callback(null, { status: 1, data: JSON.stringify(newCat) })
    } catch (err) {
        callback(err, { status: -1 })
    }
    console.log('created:', newCat)
}
function handleGet(call, callback) {
    var id = call.request.data
    if (id in categories) {
        console.log('Get:', categories[id])
        callback(null, { status: 1, data: JSON.stringify(categories[id]) })
    } else {
        console.log('Get Not found')
        callback(null, { status: 0, data: 'Category not found' })
    }
}
function handleUpdate(call, callback) {
    var cat
    try {
        cat = JSON.parse(call.request.data)
        if (cat.id in categories) {
            console.log('before Update:', JSON.stringify(categories))
            categories[cat.id] = cat
            console.log('after Update:', JSON.stringify(categories))
            callback(null, { status: 1, data: JSON.stringify(cat) })
        } else {
            console.log('Update Not found')
            callback(null, { status: 0, data: 'Category not found' })
        }
    } catch (err) {
        callback(err, { status: -1 })
    }
}
function handleDelete(call, callback) {
    var id = call.request.data
    if (id in categories) {
        console.log('before delete:', JSON.stringify(categories))
        delete categories[id]
        console.log('after delete:', JSON.stringify(categories))
        callback(null, { status: 1, data: '' })
    } else {
        console.log('Delete Not found')
        callback(null, { status: 0, data: 'Category not found' })
    }
}

module.exports = {
    create: handleCreate,
    get: handleGet,
    update: handleUpdate,
    delete: handleDelete
}
