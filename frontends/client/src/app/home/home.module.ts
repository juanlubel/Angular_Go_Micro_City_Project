import {NgModule} from '@angular/core';
import {HomeComponent} from './home.component';
import {HomeRoutingModule} from './home-routing.module';
import {CommonModule} from '@angular/common';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

@NgModule({
  declarations: [HomeComponent],
  imports: [
    HomeRoutingModule,
    // FormsModule,
    // ReactiveFormsModule,
    MatFormFieldModule,
    // MatInputModule,
    BrowserAnimationsModule,
    CommonModule,
    MatButtonModule,
    MatInputModule,

  ]
})
export class HomeModule {

}
