import os

# read file '0x0a3f9678d6b631386c2dd3de8809b48b0d1bbd56.sol.ast.json'

with open('0x0a3f9678d6b631386c2dd3de8809b48b0d1bbd56.sol.ast.json', 'r') as f:
    data = f.read()

nodetypes = []

for line in data.split('\n'):
    if '"nodeType"' in line:
        nt=line.split('"nodeType": "')[1].split('"')[0]
        if nt not in nodetypes:
            nodetypes.append(nt)

print(nodetypes)
        

# ['SourceUnit', 'PragmaDirective', 'ContractDefinition', 'Block', 'Identifier', 'BinaryOperation', 'Literal', 'IfStatement', 'Return', 'VariableDeclaration', 'ElementaryTypeName', 'VariableDeclarationStatement', 'FunctionCall', 'ExpressionStatement', 'FunctionDefinition', 'ParameterList', 'EventDefinition', 'Assignment', 'MemberAccess', 'PlaceholderStatement', 'ModifierDefinition', 'ElementaryTypeNameExpression', 'ModifierInvocation', 'UserDefinedTypeName', 'InheritanceSpecifier', 'UnaryOperation', 'Mapping', 'StructDefinition', 'UsingForDirective', 'IndexAccess', 'EnumValue', 'EnumDefinition', 'ArrayTypeName', 'ForStatement', 'Break', 'TupleExpression', 'NewExpression']