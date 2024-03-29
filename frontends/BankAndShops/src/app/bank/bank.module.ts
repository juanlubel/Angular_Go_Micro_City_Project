import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { BanksComponent } from './banks.component';
import { BankRoutingModule } from './bank-routing.module';
import { BankResolver } from './bank.resolver';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';
import { MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule,
  MatFormFieldModule,MatInputModule } from  '@angular/material';



@NgModule({
  declarations: [ BanksComponent],
  imports: [CommonModule,BankRoutingModule,FormsModule, ReactiveFormsModule,
    MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule,
    MatFormFieldModule,MatInputModule],
  providers: [
    BankResolver
  ],
  exports:[BanksComponent]
})
export class BankModule {}