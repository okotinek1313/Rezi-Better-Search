import 'dart:convert';
import 'dart:io';
import 'package:flutter/material.dart';
import 'package:window_manager/window_manager.dart';
import 'package:frontend/pages/search_page.dart';

Process? backendProcess;

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  
  // Must initialize window manager to intercept the close button
  await windowManager.ensureInitialized();

  // Cleanup any orphaned instances from previous crashes
  await Process.run('taskkill', ['/IM', 'backend.exe', '/F']);

  backendProcess = await Process.start(
    './resources/backend.exe',
    [],
    mode: ProcessStartMode.normal,
  );

  final configFile = File('./resources/shared/config.json');

  // Wait for the backend to write the port to the config file
  while (true) {
    if (configFile.existsSync()) {
      try {
        final contents = configFile.readAsStringSync();
        final config = jsonDecode(contents);
        if (config['port'] != null && config['port'] != 0) break;
      } catch (_) {}
    }
    await Future.delayed(const Duration(milliseconds: 100));
  }

  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> with WindowListener {
  
  @override
  void initState() {
    super.initState();
    windowManager.addListener(this);
    // This allows us to run code BEFORE the window actually disappears
    windowManager.setPreventClose(true);
  }

  @override
  void dispose() {
    windowManager.removeListener(this);
    super.dispose();
  }

  @override
  void onWindowClose() async {
    if (backendProcess != null) {
      // 1. Send the internal kill signal (fastest)
      backendProcess!.kill();
      
      // 2. Fire-and-forget the OS taskkill (don't 'await' this)
      // This ensures the backend dies even if it was hung, 
      // but doesn't make the user wait for the command to finish.
      Process.run('taskkill', ['/PID', '${backendProcess!.pid}', '/F', '/T']);
    }

    // 3. Close the GUI immediately
    await windowManager.destroy();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        brightness: Brightness.dark,
        textTheme: const TextTheme(
          bodyMedium: TextStyle(color: Colors.white),
        ),
      ),
      home: const SearchPage(),
    );
  }
}
