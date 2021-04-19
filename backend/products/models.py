from django.db import models


class Product(models.Model):
    id = models.CharField(max_length=20, primary_key=True)
    name = models.CharField(max_length=200)
    link = models.CharField(max_length=200)
    price = models.IntegerField()  # do przedyskutowania

    def __str__(self):
        return self.name