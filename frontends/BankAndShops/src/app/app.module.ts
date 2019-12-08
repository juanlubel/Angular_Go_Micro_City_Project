import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppRoutingModule } from './app-routing.module';
import { MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule,
         MatButtonModule,MatInputModule, MatDialogModule} from  '@angular/material';
import { HomeModule } from './home/home.module';
import {CoreModule} from './core/core.module'


import { AppComponent } from './app.component';
import {NavComponent} from './public/layout'
import {FooterComponent} from './public/layout';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';
import { CourseDialogComponent } from './public/components/dialog.component';

@NgModule({
  declarations: [AppComponent,NavComponent,FooterComponent,CourseDialogComponent],//components
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
    ReactiveFormsModule, 
    MatDialogModule 
   ],//modules
  providers: [],
  bootstrap: [AppComponent],
  entryComponents: [CourseDialogComponent]  
})
export class AppModule {}