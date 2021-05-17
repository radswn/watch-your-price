import hashlib
from django.db import models
from django.utils.translation import ugettext_lazy as _
from users.models import CustomUser as User


class Product(models.Model):
    id = models.CharField(max_length=32, editable=False, primary_key=True)
    name = models.CharField(_('name'), max_length=200)
    owner = models.ForeignKey(User, on_delete=models.CASCADE)
    link = models.URLField()
    price = models.FloatField(_('price'), )

    def save(self, **kwargs):
        if not self.id:
            self.id = hashlib.sha256(bytes(str(self.link), "utf-8")).hexdigest()
            # print(self.id)
        super().save(*kwargs)

    def __str__(self):
        return self.name

    class Meta:
        verbose_name = _('Product')
        verbose_name_plural = _('Products')
