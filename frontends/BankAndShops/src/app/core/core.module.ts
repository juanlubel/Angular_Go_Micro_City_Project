import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
/* import { HttpTokenInterceptor } from './interceptors/http.token.interceptor'; */

import {ApiService,BankService,ForumService} from './services';

@NgModule({
  imports: [
    CommonModule
  ],
  providers: [
    /* { provide: HTTP_INTERCEPTORS, useClass: HttpTokenInterceptor, multi: true }, */
    ApiService,BankService,ForumService
  ],
  declarations: []
})
export class CoreModule { }
