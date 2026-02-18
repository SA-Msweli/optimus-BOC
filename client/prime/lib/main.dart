import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});
  @override
  Widget build(BuildContext c) => MaterialApp(home: const HomePage());
}

class HomePage extends StatefulWidget {
  const HomePage({super.key});
  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  String status = 'checking...';
  @override
  void initState() {
    super.initState();
    check();
  }

  Future<void> check() async {
    final r = await http.get(
      Uri.parse('http://10.0.2.2:8080/health'),
    ); // emulator host
    setState(() => status = (r.statusCode == 200) ? 'OK' : 'ERR');
  }

  @override
  Widget build(BuildContext c) => Scaffold(
    appBar: AppBar(title: const Text('Optimus (dev)')),
    body: Center(child: Text('backend health: $status')),
    floatingActionButton: FloatingActionButton(
      onPressed: check,
      child: const Icon(Icons.refresh),
    ),
  );
}
