import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { MatListOption } from '@angular/material/list';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ItemProps } from './home.model';
import { HomeService } from './home.service';
import { ListUtil } from 'src/app/shared/utils/list.util';

type SortKeys = keyof Omit<ItemProps, 'id'>;

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  items: ItemProps[] = [];
  selectedItems: ItemProps[] = [];
  nameControl = new FormControl('',[Validators.required, Validators.minLength(1)]);
  formSubmited = false;
  sortLabels: Record<SortKeys, string> = {
    name: 'Nome',
    createdAt: 'Data de criação',
  };
  sortedBy: SortKeys = 'createdAt';
  sortedReverse = false;
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
    this.homeService.getItems().subscribe((res) => {
      this.items = res ?? [];
      this.sort(this.sortedBy);
      this.loading = false;
    });
  }

  createItem() {
    this.formSubmited = true;
    if (this.nameControl.valid) {
      this.loading = true;
      let item: ItemProps = {
        name: `${this.nameControl.value}`,
        createdAt: new Date(),
      };

      this.homeService.insertItem(item).subscribe(
        () => {
          this.snackBar.open('Item inserido com sucesso', 'OK', {
            panelClass: ['snack-success'],
            duration: 1000
          });
          this.loadItems();
          this.nameControl.reset();
          this.loading = this.formSubmited = false;
        },
        () => {
          this.snackBar.open('Erro ao inserir item', 'OK', {
            panelClass: ['snack-fail'],
            duration: 1000
          });
        }
      );
    } else {
      this.nameControl.updateValueAndValidity();
    }
  }

  removeItems() {
    this.selectedItems.forEach(
      (selected) => {
        this.loading = true;
        this.homeService.deleteItem(selected.id!).subscribe(() => {
          this.loadItems();
        });
      }
    );
    this.selectedItems = [];
  }

  sort(key: string) {
    this.loading = true;
    this.sortedBy = key as SortKeys;
    this.items = ListUtil.toSorted<ItemProps>(this.items, this.sortedBy, this.sortedReverse);
    this.loading = false;
  }

  reverseSort() {
    this.sortedReverse = !this.sortedReverse;
    this.sort(this.sortedBy);
  }

  onSelection(selectedItems: MatListOption[]) {
    this.selectedItems = selectedItems.map((item) => item.value);
  }
}
