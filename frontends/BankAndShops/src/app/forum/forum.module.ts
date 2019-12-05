import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';
import { MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule,
  MatFormFieldModule,MatInputModule,MatGridListModule,MatDividerModule,MatCardModule } from  '@angular/material';
import { ForumComponent } from './forum.component';
import { ForumRoutingModule } from './forum-routing.module';



@NgModule({
  declarations: [ForumComponent],
  imports: [CommonModule,FormsModule, ReactiveFormsModule,
    ForumRoutingModule,
    MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule,
    MatFormFieldModule,MatInputModule,MatGridListModule,MatDividerModule,MatCardModule],
  providers: [
    
  ],
  exports:[]
})
export class ForumModule {}