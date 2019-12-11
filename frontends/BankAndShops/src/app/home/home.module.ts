import { NgModule } from "@angular/core";


import {HomeRoutingModule  } from "./home-routing.module";
import { HomeComponent } from "./home.component";
import { FormsModule, ReactiveFormsModule} from '@angular/forms';
import { MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule,
        MatFormFieldModule,MatInputModule } from  '@angular/material';


@NgModule({
  imports: [ HomeRoutingModule,FormsModule, ReactiveFormsModule,
             MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule,
             MatFormFieldModule,MatInputModule],
  declarations: [HomeComponent]
})
export class HomeModule {}