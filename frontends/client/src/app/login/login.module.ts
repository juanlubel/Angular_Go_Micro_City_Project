import {NgModule} from '@angular/core';

import {LoginComponent} from './login.component';
import {LoginRoutingModule} from './login-routing.module';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {CommonModule} from '@angular/common';


@NgModule({
  declarations: [LoginComponent],
  imports: [
    LoginRoutingModule,
    // FormsModule,
    // ReactiveFormsModule,
    // MatFormFieldModule,
    // MatInputModule,
    CommonModule
  ]
})
export class LoginModule {
}
