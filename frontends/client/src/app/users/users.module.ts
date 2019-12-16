import {NgModule} from '@angular/core';

import {UsersComponent} from './users.component';
import {UsersRoutingModule} from './users-routing.module';
import {UsersListComponent} from './users-list/users-list.component';
import {CommonModule} from '@angular/common';
import {UserDetailsComponent} from './user-details/user-details.component';
import {FormsModule} from "@angular/forms";


@NgModule({
    declarations: [UsersComponent, UsersListComponent, UserDetailsComponent],
  imports: [
    UsersRoutingModule,
    CommonModule,
    FormsModule
  ],
  exports: []
})
export class UsersModule {
}
