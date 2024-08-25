import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { MatListOption } from '@angular/material/list';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ItemProps } from './home.model';
import { HomeService } from './home.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})

export class HomeComponent implements OnInit{
  selectedItems: ItemProps[] = [];
  items: ItemProps[] = [];
  nameControl = new FormControl('',[Validators.required, Validators.minLength(1)]);
  formSubmitedd = false;
  sortLabels = {
    name: 'Nome',
    createdAt: 'Data de criação'
  }
  filtersMap = {
    none: 'Nenhum filtro',
  }
  selectedFilter = 'none';
  sortedBy = {
    key: 'database',
    reverse: false
  }
  loading = false;

  constructor(
    private homeService: HomeService,
    private snackBar: MatSnackBar
  ){ }

  ngOnInit(): void {
    this.loadItems();
  }

  loadItems() {
    this.loading = true;
    this.homeService.getItems().subscribe({
      next: (res) => {
        this.items = res;
        this.sort(this.sortedBy.key); 
        this.loading = false;
      },
      error: () => {
        this.loading = false;
      }
    });
  }

  isNameInvalid() {
    return !this.nameControl.valid && this.formSubmitedd;
  }

  createItem(){
    this.formSubmitedd = true;
    if(!this.isNameInvalid()){
      this.loading = true;
      let item: ItemProps = {
        name: `${this.nameControl.value}`,
        createdAt: new Date(),
      };

      this.homeService.insertItem(item).subscribe({
        next: () => {
          this.snackBar.open('Item inserido com sucesso', 'OK', {
            panelClass: ['snack-success'],
            duration: 1000
          });
          this.loadItems();
          this.nameControl.reset();
          this.loading = this.formSubmitedd = false;
        },
        error: () => {
          this.snackBar.open('Erro ao inserir item', 'OK', {
            panelClass: ['snack-fail'],
            duration: 1000
          });
          this.loading = this.formSubmitedd = false;
        }
      });
    } else {
      this.nameControl.updateValueAndValidity();
    }
  }

  removeItems(){
    this.selectedItems.forEach(
      (selected) => {
        this.loading = true;
        this.homeService.deleteItem(selected.id!).subscribe(() => {
          this.loadItems();
          this.loading = false;
        });
      }
    );
    this.selectedItems = [];
  }

  sort(key: string){
    this.loading = true;
    this.sortedBy.key = key;
    if(this.sortedBy.reverse){
      this.items = this.items.sort((a: any, b: any) => {
        if (a[key] > b[key]) {
          return -1;
        } else if (a[key] < b[key]) {
            return 1;
        }
        return 0;
      });
    }else{
      this.items = this.items.sort((a: any, b: any) => {
        if (a[key] > b[key]) {
          return 1;
        } else if (a[key] < b[key]) {
            return -1;
        }
        return 0;
      });
    }
    this.loading = false;
  }

  reverseSort(){
    this.sortedBy.reverse = !this.sortedBy.reverse;
    this.sort(this.sortedBy.key);
  }

  filter(key: any, value?: any){
    this.loading = true;
    this.selectedFilter = key != 'database'? key: value;
    this.loadItems();
    if(key != 'none'){
      this.items = this.items.filter((item) => item[key] == value);
    }
    this.loading = false;
  }

  onSelection(event: any, selectedItems: MatListOption[]){
    this.selectedItems = selectedItems.map((item) => item.value);
  }
}
