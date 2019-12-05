import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppRoutingModule } from './app-routing.module';
import { MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule,MatInputModule } from  '@angular/material';
import { HomeModule } from './home/home.module';
import {CoreModule} from './core/core.module'


import { AppComponent } from './app.component';
import {NavComponent} from './public/layout'
import {FooterComponent} from './public/layout';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

@NgModule({
  declarations: [AppComponent,NavComponent,FooterComponent],//components
  imports: [BrowserModule, AppRoutingModule,BrowserAnimationsModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
    MatButtonModule,
    MatIconModule,
    MatInputModule,
    HomeModule,
    CoreModule,
    HttpClientModule,
    FormsModule, 
    ReactiveFormsModule
   ],//modules
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}