import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { BanksComponent } from './banks.component';
import { BankRoutingModule } from './bank-routing.module';



@NgModule({
  declarations: [ BanksComponent],
  imports: [CommonModule,BankRoutingModule],
  providers: []
})
export class BankModule {}