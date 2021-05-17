from django.contrib import admin
from .models import Product
from django.utils.translation import gettext_lazy as _
# Register your models here.
admin.site.register(Product)
