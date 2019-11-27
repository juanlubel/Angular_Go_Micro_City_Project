import {NgModule} from '@angular/core';

import {UsersComponent} from './users.component';
import {UsersRoutingModule} from './users-routing.module';
import {UsersListComponent} from './users-list/users-list.component';


@NgModule({
  declarations: [UsersComponent, UsersListComponent],
  imports: [
    UsersRoutingModule
  ],
  exports: []
})
export class UsersModule {
}
