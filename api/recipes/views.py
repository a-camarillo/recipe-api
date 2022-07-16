from recipes.models import Recipe, Ingredient
from recipes.serializers import RecipeSerializer, IngredientSerializer, RecipePostSerializer
from rest_framework import generics
from rest_framework.decorators import api_view
from rest_framework.response import Response
from rest_framework.reverse import reverse

@api_view(['GET'])
def api_root(request, format=None):
    return Response({
        'recipes': reverse('recipe-list', request=request, format=format),
        'ingredients': reverse('ingredient-list', request=request, format=format)
    })


class RecipeList(generics.ListCreateAPIView):
    queryset = Recipe.objects.all()

    def list(self, request):
        queryset = self.get_queryset() 
        
        #creating serializer instance to specify fields, context arg required for HyperLinkIdentityField
        serializer = RecipeSerializer(queryset, many=True, fields=('url','recipe_name'), context={'request':request})

        return Response(serializer.data)

    def get_serializer_class(self):
        if self.request.method == 'GET':
            return RecipeSerializer
        if self.request.method == 'POST':
            return RecipePostSerializer


class IngredientList(generics.ListAPIView):
    queryset = Ingredient.objects.all()
    serializer_class = IngredientSerializer

    def list(self, request):
        queryset = self.get_queryset()

        serializer = IngredientSerializer(queryset, many=True, fields=('url','ingredient_name'), context={'request':request})

        return Response(serializer.data)

class RecipeDetail(generics.RetrieveUpdateDestroyAPIView):
    queryset = Recipe.objects.all()
    serializer_class = RecipeSerializer
    
class IngredientDetail(generics.RetrieveUpdateDestroyAPIView):
    queryset = Ingredient.objects.all()
    serializer_class = IngredientSerializer