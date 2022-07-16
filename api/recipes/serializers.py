from rest_framework import serializers
from recipes.models import Recipe, Ingredient

class CustomFieldSerializer(serializers.ModelSerializer):
    '''
    Custom Serializer class to allow for selection of fields to be displayed in views
    '''
    def __init__(self, *args, **kwargs):
        
        fields = kwargs.pop('fields', None)

        super().__init__(*args,**kwargs)

        if fields is not None:
            allowed = set(fields) #all fields being declared
            existing = set(self.fields) #all fields for the serializer
            for field_name in existing - allowed: #field names that are not being specified
                self.fields.pop(field_name)


class RecipeSerializer(serializers.HyperlinkedModelSerializer,
                       CustomFieldSerializer):

    ingredients = serializers.HyperlinkedRelatedField(
        many = True,
        view_name = 'ingredient-detail',
        read_only = True
        )    
    class Meta:
        model = Recipe
        fields = ['url','id','recipe_name','ingredients','instruction']

class IngredientSerializer(serializers.HyperlinkedModelSerializer,
                           CustomFieldSerializer):
    class Meta:
        model = Ingredient
        fields = ['url','id','ingredient_name']

class IngredientPostSerializer(serializers.ModelSerializer):
    class Meta:
        model = Ingredient
        fields = ['ingredient_name']

class RecipePostSerializer(serializers.ModelSerializer):
    ingredients = IngredientPostSerializer(many=True)

    def create(self, validated_data):
        ingredients_data = validated_data.pop('ingredients')
        recipe = Recipe.objects.create(**validated_data)
        for ingredient_data in ingredients_data:
            Ingredient.objects.update_or_create(recipe=recipe, **ingredient_data)
        return recipe

    class Meta:
        model = Recipe
        fields = ['recipe_name','ingredients','instruction']