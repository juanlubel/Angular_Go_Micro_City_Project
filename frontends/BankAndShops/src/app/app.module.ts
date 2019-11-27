import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppRoutingModule } from './app-routing.module';
import { MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule } from  '@angular/material';
import { HomeModule } from './home/home.module';
import {CoreModule} from './core/core.module'
import { ReactiveFormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import {NavComponent} from './public/layout'
import {FooterComponent} from './public/layout';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [AppComponent,NavComponent,FooterComponent],//components
  imports: [BrowserModule, AppRoutingModule,BrowserAnimationsModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
    MatButtonModule,
    MatIconModule,
    HomeModule,
    CoreModule,
    HttpClientModule,
    ReactiveFormsModule],//modules
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}