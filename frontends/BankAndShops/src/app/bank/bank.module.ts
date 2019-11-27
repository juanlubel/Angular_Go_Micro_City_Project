import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { BanksComponent } from './banks.component';
import { BankRoutingModule } from './bank-routing.module';
import { BankResolver } from './bank.resolver';



@NgModule({
  declarations: [ BanksComponent],
  imports: [CommonModule,BankRoutingModule],
  providers: [
    BankResolver
  ],
  exports:[BanksComponent]
})
export class BankModule {}