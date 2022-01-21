/*
|--------------------------------------------------------------------------
| Routes
|--------------------------------------------------------------------------
|
| This file is dedicated for defining HTTP routes. A single file is enough
| for majority of projects, however you can define routes in different
| files and just make sure to import them inside this file. For example
|
| Define routes in following two files
| ├── start/routes/cart.ts
| ├── start/routes/customer.ts
|
| and then import them inside `start/routes.ts` as follows
|
| import './routes/cart'
| import './routes/customer'
|
*/

import Route from '@ioc:Adonis/Core/Route'

Route.group(() => {
  // Messages.
  Route.group(() => {
    Route.delete('/', 'MessagesController.delete').middleware(['auth'])
    
    Route.get('/', 'MessagesController.getMultiple')
    Route.post('/', 'MessagesController.insert')
    Route.post('/ack', 'MessagesController.ack')
  }).prefix('messages')

  // Clients.
  Route.group(() => {
    Route.group(() => {
      Route.get('/:id', 'ClientsController.get')
      Route.get('/', 'ClientsController.getMultiple')
      
      Route.group(() => {
        Route.get('/count', 'ClientsController.getCount')
        Route.get('/demographic', 'ClientsController.demographic')
      }).prefix('stats')
    }).middleware(['auth'])
      
    Route.post('/', 'ClientsController.insert')

  }).prefix('clients')
}).prefix('v1')