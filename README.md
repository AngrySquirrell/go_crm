# Mini-CRM CLI

Un gestionnaire de contacts simple et efficace en ligne de commande, écrit en Go.

## Fonctionnalités

- **CRUD complet** : Ajouter, lister, mettre à jour et supprimer des contacts
- **Interface CLI professionnelle** avec Cobra
- **Configuration externe** avec Viper
- **Multiples backends de stockage** :
  - GORM/SQLite (base de données)
  - Fichier JSON (persistance simple)
  - Mémoire (stockage éphémère)

## Installation

Clonez le repository et compilez le projet :

```bash
go build
```

## Configuration

Le comportement de l'application est contrôlé par le fichier `config.yaml` :

```yaml
storage:
  type: "json" # Options: gorm, json, memory
  database_path: "crm.db"
  json_path: "contacts.json"
```

## Utilisation

### Commandes disponibles

#### Ajouter un contact

```bash
./go_crm add --name "Jean Dupont" --email "jean@example.com"
```

#### Lister tous les contacts

```bash
./go_crm list
```

#### Mettre à jour un contact

```bash
./go_crm update [id] --name "Nouveau nom" --email "nouveau@example.com"
```

#### Supprimer un contact

```bash
./go_crm delete [id]
```

#### Aide

```bash
./go_crm --help
./go_crm [commande] --help
```

## Versions précédentes

-> [v1](https://github.com/AngrySquirrell/go/tree/td1)

-> [v1.1](https://github.com/AngrySquirrell/go/tree/td1.1)

-> [v1.2](https://github.com/AngrySquirrell/go/tree/td1.2)
