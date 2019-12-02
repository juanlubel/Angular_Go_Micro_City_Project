import {NgModule} from '@angular/core';

import {UsersComponent} from './users.component';
import {UsersRoutingModule} from './users-routing.module';
import {UsersListComponent} from './users-list/users-list.component';
import {CommonModule} from '@angular/common';


@NgModule({
  declarations: [UsersComponent, UsersListComponent],
  imports: [
    UsersRoutingModule,
    CommonModule
  ],
  exports: []
})
export class UsersModule {
}
